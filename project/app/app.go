package main

import (
  "github.com/go-redis/redis/v8"
  "net/http"
  "context"
)

var (
  rdb *redis.Client
  ctx = context.Background()
)

func main() {
  // Соединение с Redis
  rdb = redis.NewClient(&redis.Options{
    Addr:     "127.0.0.1:6379", // адрес сервера Redis
    Password: "redis_pass", // пароль (если есть)
    DB:       0,  // используемая база данных
  })

  http.HandleFunc("/set_key", setHandler)
  http.HandleFunc("/get_key", getHandler)
  http.HandleFunc("/del_key", deleteHandler)

  // Запуск сервера
  http.ListenAndServe(":8089", nil)
}

func setHandler(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    value := r.URL.Query().Get("value")

    err := rdb.Set(ctx, key, value, 0).Err()
    if err != nil {
        http.Error(w, "Unable to set value", http.StatusInternalServerError)
        return
    }

    w.Write([]byte("Value was set successfully"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")

    value, err := rdb.Get(ctx, key).Result()
    if err != nil {
        http.Error(w, "Unable to get value", http.StatusInternalServerError)
        return
    }

    w.Write([]byte(value))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")

    err := rdb.Del(ctx, key).Err()
    if err != nil {
        http.Error(w, "Unable to delete key", http.StatusInternalServerError)
        return
    }

    w.Write([]byte("Key deleted successfully"))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
   correctURI := "/set_key?key=<key>&value<key_value>"
   if r.URL.Path != correctURI {
      http.Error(w, "Access denied", http.StatusForbidden)
      return
   }
   // Ваш код обработки для корректного URI здесь
}
