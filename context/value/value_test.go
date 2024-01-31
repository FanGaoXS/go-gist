package value

import (
	"context"
	"testing"
)

func setInCtx(ctx context.Context, key string, value any) context.Context {
	newCtx := context.WithValue(ctx, key, value)
	return newCtx
}

func getFromCtx(ctx context.Context, key string) any {
	return ctx.Value(key)
}

func TestValue(t *testing.T) {
	k := "k"
	v := "v"
	ctx := setInCtx(context.Background(), k, v)
	got := getFromCtx(ctx, k)
	if got != v {
		t.Errorf("got shoule be %s\n", v)
	}

	type User struct {
		Name string
		Age  int
	}

	k = "user::context"
	u := User{
		Name: "foo",
		Age:  10,
	}
	ctx = setInCtx(context.Background(), k, u)
	got = getFromCtx(ctx, k)
	if got == nil {
		t.Errorf("got should not be nil")
	}
	gotU := got.(User) // 判定got的类型是User
	if gotU.Name != u.Name && gotU.Age != u.Age {
		t.Errorf("got shoule be the same as %v\n", u)
	}
}
