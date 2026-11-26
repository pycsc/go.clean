/* ************************************************************************** */
/*                                                                            */
/*                                                            +#              */
/*   delete_test.go                                         +##               */
/*                                                         +###+              */
/*   By: lucas <lucas@tsiaraso.com>                      +########+           */
/*                                                           +###+            */
/*   Created: 2026/03/12 12:48:56 by lucas                   ##+              */
/*   Updated: 2026/11/26 14:00:36 by lucas                  #+                */
/*                                                                            */
/* ************************************************************************** */

package usecases_test

import "testing"

func TestDelete(t *testing.T) {
	uc := newUC()
	uc.Create("to delete")
	if !uc.Delete(1) {
		t.Fatal("expected delete ok")
	}
	if uc.Delete(1) {
		t.Fatal("expected delete fail for missing")
	}
}
