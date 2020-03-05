package operation

import "errors"

const(
    SUM = "sum"
    SUB = "sub"
    MUL = "mul"
    DIV = "div"
)

type Operation struct {
    A float64
    B float64
    Operation string
}

func (op Operation) Calculate() (float64, error) {
    if op.Operation == DIV && op.B == 0 {
        return 0, errors.New("Division by 0")
    }
    
    switch op.Operation {
    case SUM: return op.A + op.B, nil
    case SUB: return op.A - op.B, nil
    case MUL: return op.A * op.B, nil
    case DIV: return op.A / op.B, nil
    default: return 0, errors.New("Invalid operation")
    }
}
