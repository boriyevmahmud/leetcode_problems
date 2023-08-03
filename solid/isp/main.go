package main

// type Printer interface {
//     PrintDocument(document string)
//     ScanDocument() string
//     FaxDocument(document string)
//     CopyDocument() string
// }

type Printer interface {
    PrintDocument(document string)
}

type Scanner interface {
    ScanDocument() string
}

type FaxMachine interface {
    FaxDocument(document string)
}

type Copier interface {
    CopyDocument() string
}

func PrintSomething(p Printer) {
    p.PrintDocument("Hello, world!")
}
