/* ************************************************************************** */
/*                                                                            */
/*                                                            +#              */
/*   helpers_test.go                                        +##               */
/*                                                         +###+              */
/*   By: lucas <lucas@tsiaraso.com>                      +########+           */
/*                                                           +###+            */
/*   Created: 2026/03/12 12:48:56 by lucas                   ##+              */
/*   Updated: 2026/11/26 14:00:36 by lucas                  #+                */
/*                                                                            */
/* ************************************************************************** */

package usecases_test

import (
	"go.clean/internal/infrastructure"
	"go.clean/internal/usecases"
)

func newUC() *usecases.TodoUseCase {
	return usecases.NewUseCase(infrastructure.NewMemoryRepo())
}
