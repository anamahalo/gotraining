package main

import "fmt"

type reader interface {
    read(b []byte) (int, error)
}

type file struct {
    name string
}

// receiver here declared w/o var name
//  common practice when method doesn't need to access
//  anything from receiver value
func (file) read(b []byte) (int, error) {
    s := "<rss><channel><title>Golang</title></channel></rss>"
    copy(b, s)
    return len(s), nil
}

type pipe struct {
    name string
}

// receiver here declared w/o var name
//  common practice when method doesn't need to access
//  anything from receiver value
func (pipe) read(b []byte) (int, error) {
    s := `{name: "Jane", title: developer}`
    copy(b, s)
    return len(s), nil
}

// .: The concrete type file and pipe now
//    implement the reader interface using
//    value receivers
//    (Behavior data can exhibit is defined
//     by semantic used:
//     https://www.ardanlabs.com/blog/2017/06/design-philosophy-on-data-and-semantics.html)

// .: Now that methods are declared using value receivers,
//    values and pointers of these concrete types
//    can be passsed into the polymorphic func retrieve()

func main() {
    f := file{"data.json"}
    p := pipe{"cfg_service"}

    retrieve(f)
    retrieve(p)
}

// can accept any value or pointer
// that implements reader interface
// .: Decoupling achieved
//      Reiteration:
//          func retriever() not asking for a reader value
//          b/c it is valueless - does not exist
//          func retriever() is asking for concrete data
//          that exhibits the correct BEHAVIORS
func retrieve(r reader) error {
    data := make([]byte, 100)

    len, err := r.read(data)
    if err != nil {
        return err
    }

    fmt.Println(string(data[:len]))
    return nil
}
