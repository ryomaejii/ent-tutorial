package todo

import (
	"context"
	"fmt"
	"log"
	"todo/ent"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func Todo() {
    // Create an ent.Client with in-memory SQLite database.
    client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }
    defer client.Close()
    ctx := context.Background()
    // Run the automatic migration tool to create all schema resources.
    if err := client.Schema.Create(ctx); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
    task1, err := client.Todo.Create().SetText("Add GraphQL Example").Save(ctx)
    if err != nil {
        log.Fatalf("failed creating a todo: %v", err)
    }
    fmt.Printf("%d: %q\n", task1.ID, task1.Text)
    task2, err := client.Todo.Create().SetText("Add Tracing Example").Save(ctx)
    if err != nil {
        log.Fatalf("failed creating a todo: %v", err)
    }
    if err := task2.Update().SetParent(task1).Exec(ctx); err != nil {
        log.Fatalf("failed connecting todo2 to its parent: %v", err)
    }
    fmt.Printf("%d: %q\n", task2.ID, task2.Text)
    // Output:
    // 1: "Add GraphQL Example"
    // 2: "Add Tracing Example"
}