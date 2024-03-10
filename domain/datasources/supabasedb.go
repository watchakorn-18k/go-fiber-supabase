package datasources

import (
	"fmt"
	"os"

	"github.com/supabase-community/supabase-go"
)

type SupabaseDB struct {
	Client *supabase.Client
}

type ISupabaseDB interface {
}

func NewSupabaseDB() *SupabaseDB {

	client, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"), nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}

	return &SupabaseDB{
		Client: client,
	}
}
