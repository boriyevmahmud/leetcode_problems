package main

import (
    "crypto/sha1"
    "encoding/hex"
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    downloadedHash, err := convertToSHA1("201c70a6-a621-48bc-ac41-3ad87bb72949")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Download the image and save it to a file
    if err := saveDownloadedImage("201c70a6-a621-48bc-ac41-3ad87bb72949"); err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Image downloaded successfully with SHA-1 hash: %s\n", downloadedHash)
}

func calculateSHA1Hash(data []byte) string {
    hash := sha1.New()
    hash.Write(data)
    hashBytes := hash.Sum(nil)
    hashString := hex.EncodeToString(hashBytes)
    return hashString
}

func convertToSHA1(img string) (string, error) {
    downloadURL := "https://test.cdn.delever.uz/delever/" + img

    resp, err := http.Get(downloadURL)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    hash := calculateSHA1Hash(data)
    return hash, nil
}

func saveDownloadedImage(img string) error {
    downloadURL := "https://test.cdn.delever.uz/delever/" + img

    resp, err := http.Get(downloadURL)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Create a new file for the downloaded image
    file, err := os.Create("downloaded_image.jpg") // Change the file extension based on the image format
    if err != nil {
        return err
    }
    defer file.Close()

    // Copy the downloaded data to the file
    _, err = io.Copy(file, resp.Body)
    if err != nil {
        return err
    }

    return nil
}
