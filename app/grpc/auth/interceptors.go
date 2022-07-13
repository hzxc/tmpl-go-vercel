package auth

import (
	"context"
	"crypto/rsa"
	"strings"
	"tmpl-go-vercel/app/shared/auth/token"

	wrapper_stream "tmpl-go-vercel/app/grpc/wrapper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
	bearerPrefix        = "Bearer "
)

// Interceptor creates a grpc auth interceptor.
func Interceptor(pubKey *rsa.PublicKey) *interceptor {
	i := &interceptor{
		verifier: &token.JWTTokenVerifier{
			PublicKey: pubKey,
		},
	}
	return i
}

type tokenVerifier interface {
	Verify(token string) (string, error)
}

type interceptor struct {
	verifier tokenVerifier
}

// AccountID defines account id object.
type AccountID string

func (a AccountID) String() string {
	return string(a)
}

type accountIDKey struct{}

func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	aid, err := i.verifier.Verify(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token not valid: %v", err)
	}

	return handler(ContextWithAccountID(ctx, AccountID(aid)), req)
}

func tokenFromContext(c context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", unauthenticated
	}

	tkn := ""
	for _, v := range m[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			tkn = v[len(bearerPrefix):]
		}
	}
	if tkn == "" {
		return "", unauthenticated
	}

	return tkn, nil
}

// ContextWithAccountID creates a context with given account ID.
func ContextWithAccountID(c context.Context, aid AccountID) context.Context {
	return context.WithValue(c, accountIDKey{}, aid)
}

// AccountIDFromContext gets account id from context.
// Returns unauthenticated error if no account id is available.
func AccountIDFromContext(c context.Context) (AccountID, error) {
	v := c.Value(accountIDKey{})
	aid, ok := v.(AccountID)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return aid, nil
}

// UnaryServerInterceptor returns a new unary server interceptor for OpenTracing.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		aid, err := auth(ctx, o)
		if err != nil {
			// return nil, err
			return nil, status.Error(codes.Unauthenticated, "token not valid")
		}

		return handler(ContextWithAccountID(ctx, AccountID(aid)), req)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor for OpenTracing.
func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := evaluateOptions(opts)
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		w := wrapper_stream.NewStreamContextWrapper(stream)

		aid, err := auth(w.Context(), o)

		if err != nil {
			return status.Error(codes.Unauthenticated, "token not valid")
		}

		ctx := ContextWithAccountID(stream.Context(), AccountID(aid))

		w.SetContext(ctx)
		return handler(srv, w)
	}
}

func auth(ctx context.Context, opts *options) (string, error) {
	i := Interceptor(opts.pubKey)
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return "", err
	}

	aid, err := i.verifier.Verify(tkn)
	if err != nil {
		return "", err
	}

	return aid, nil
}
