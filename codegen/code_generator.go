
package codegen

import "os"

type CodeGenerator interface {
    Execute ( f * os.File ) error
}
