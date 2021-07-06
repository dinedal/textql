# Maintainer: Aniket-Pradhan aniket17133@iiitd.ac.in
# Owner/Cofntributer: Paul Bergeron https://github.com/dinedal

pkgname=textql-git
pkgver=2.0.3
pkgrel=1
pkgdesc="Execute SQL against structured text like CSV or TSV"
arch=('x86_64' 'i686')
url="https://github.com/dinedal/textql"
license=('MIT')
depends=('go')
makedepends=('git')
options=('!strip' '!emptydirs')
_gourl=github.com/dinedal/textql

build() {
  GOPATH="$srcdir" go get -v -u ${_gourl}/...
}

check() {
  echo $GOPATH
  echo $srcdir
  GOPATH="$GOPATH:$srcdir" go test -v ${_gourl}/...
}

package() {
  mkdir -p "$pkgdir/usr/bin"
  install -p -m755 "$srcdir/bin/"* "$pkgdir/usr/bin"

  mkdir -p "$pkgdir/$GOPATH"
  cp -Rv --preserve=timestamps "$srcdir/"{src,pkg} "$pkgdir/$GOPATH"

  for f in LICENSE COPYING LICENSE.* COPYING.*; do
    if [ -e "$srcdir/src/$_gourl/$f" ]; then
      install -Dm644 "$srcdir/src/$_gourl/$f" \
        "$pkgdir/usr/share/licenses/$pkgname/$f"
    fi
  done
}

# vim:set ts=2 sw=2 et:
