package repositories

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

func InitializeSupabaseClient() *supa.Client {
	// only load the .env file when running locally
	// check for a RAILWAY_ENVIRONMENT, if not found, code is running locally
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	if supabaseUrl == "" {
		log.Fatal("SUPABASE_URL environment variable is not set")
	}

	supabaseKey := os.Getenv("SUPABASE_KEY")
	if supabaseKey == "" {
		log.Fatal("SUPABASE_KEY environment variable is not set")
	}

	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	return supabase
}
