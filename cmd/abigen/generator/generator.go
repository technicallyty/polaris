// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package generator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dave/jennifer/jen"
)

// `abiPkg` is the package name of the `abi` package.
const abiPkg = "github.com/berachain/stargazer/types/abi"

// `Run` generates a golang file representing a `CompiliedContract` from
// the output of `forge build`.
func Run(packageName, inputPath, outputPath, varName string) error {
	// Read the JSON data from the file.
	data, err := os.ReadFile(inputPath) //#nosec:G304 // not part of live chain.
	if err != nil {
		return err
	}

	// Declare a variable to hold the unmarshaled JSON data.
	var forgeJSON interface{}

	// Unmarshal the JSON data into the forgeJSON variable.
	if err = json.Unmarshal(data, &forgeJSON); err != nil {
		return err
	}

	// Load ABI String.
	abiString := forgeJSON.(map[string]interface{})["abi"]
	abiStringMarshalled, err := json.Marshal(abiString)
	if err != nil {
		return err
	}

	// Load Bytecode String.
	bytecode := forgeJSON.(map[string]interface{})["bytecode"].(map[string]interface{})["object"]
	bytecodeMarshalled, err := json.Marshal(bytecode)
	if err != nil {
		return err
	}

	// Build the Golang file.
	f := buildFile(packageName, varName, string(abiStringMarshalled), string(bytecodeMarshalled))

	// Create Path.
	outFile, err := os.Create(outputPath) //#nosec:G304 // not part of live chain.
	if err != nil {
		log.Fatal(err)
	}

	// Output to file.
	_, err = outFile.WriteString(fmt.Sprintf("%#v", f) + "\n")
	if err != nil {
		return err
	}

	// Sync the file.
	if err = outFile.Sync(); err != nil {
		return err
	}

	return nil
}

// `buildFile` builds a golang file representing a `CompiliedContract` from the
// provided packageName, varName, abiStringMarshalled, and bytecodeMarshalled.
func buildFile(packageName, varName, abiStringMarshalled, bytecodeMarshalled string) *jen.File {
	f := jen.NewFilePath(packageName)
	f.HeaderComment("Code generated by stargazer. DO NOT EDIT.")
	f.Var().Id(varName).Qual(abiPkg, "CompiliedContract")
	f.Func().Id("init").Params().Block(
		jen.Id(varName).Op("=").Qual(abiPkg, "BuildCompiledContract").Call(
			jen.Id("\""+strings.ReplaceAll(abiStringMarshalled, "\"", "\\\"")+"\""),
			jen.Id(bytecodeMarshalled),
		),
	)
	return f
}
