package main

import (
	"image/png"
	"os"
	"flag"
	"fmt"
	"github.com/nfnt/resize"
)

var fileName = flag.String("f", "test.png", "filename")
var height = flag.Uint("h", 80, "height")
var width = flag.Uint("w", 80, "width")

var charSet string = "$@BB8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "

//将png图片转换为字符图片
func main() {
	flag.Parse()
	file, err := os.Open(*fileName)
	defer file.Close()
	if err != nil {
		fmt.Printf("File Open Error:%s", err.Error())
		return
	}
	img, err := png.Decode(file)
	if err != nil {
		fmt.Printf("Decode Error:%s", err.Error())
		return
	}

	result := []byte{}
	img = resize.Resize(*height, *width, img, resize.NearestNeighbor)
	for i := 0; uint(i) < *height; i++ {
		for j := 0; uint(j) < *width; j++ {
			c := getChar(img.At(i, j).RGBA())
			result = append(result, c)
		}
		result = append(result, '\n')
	}

	fmt.Printf(string(result))
}


func getChar(r, g, b, a uint32) byte {
	if a == 0 {
		return ' '
	}
	length := len(charSet)
	index := 0.2126 * float64(r) + 0.7152 * float64(g) + 0.0722 * float64(b)
	num := int(index) % length
	return charSet[num]

}


/*


                                   _bd?888,k`k?
                                qO{_>.:}'`"xrzOfWxq
                              qu+^jdd-}8}dddOc:::_"dz
                             Y')j!ddd_,!dddddX::::::px
                           qz-{)nddddfxdddddd")::::::z1J
                          q1c::p~d;kdiOBxMdddLf::::::::qa
                         qv':::!d;q^;@L@-'jddd@:::::::::XY/
                         c}::)rld-ddOvcddL;ddd@::::::::::'pp
                        nbxLdC @ddddOXjdddadd-qx::::::::::'lq
                       I|h;dddd*B{8v_-kdddddd]aqL_)::::::::}B
                     *q0l;jddddjoIq^Cr@}jdddi;dddvu~?:::::::bX
                    &Lda\ ]@;!d!!.bdQ]di@&_]^dddddB;#w:::::::&{
                   >|dddMddd U*ax.'q.oddddddd~k]&Uf-dXb::::::'lq
                   "!dddddddddad[\..>>ddd &UUiYdddddddCr::::::{]
                   -dvqqqqL-!dddM}cn@ddaq|!ddddddddddddC'::::::0q
                   -dvqqqq1mZt|ddd@ddddddddddddddddddddMB::::::v^
                   Uaddddddddddddd]dddddddddddM---qq}ddd\`:::::)W
                   lfddddddddddddd&ddddd-&_&__kqqqB-MdddBd::::::,
                    !-ddddM;UUddd-fdddddddddddddddddddddd+::::::{
                     pU&U&k-dddddqBddddqqad!BqczcccCqdddd*-:::::ZL
                     Y*U~ddddddddq-jLkL(<`.Co'......}hdddaZ:::::z;
                     iY!?_#Y_zc>?.B~[w......IL<hh&\..Uddddo:::::Xk
                     JOd#I.......................>~>.Mddddr:::::?+
                     /*d!_...........................Uddddu:::::Xk
                      |ddXq..........................?ddddz:::::xf
                      qdd!U.........................&^dddd_:::::ZB
                      bddd$.................(wjcr...}ddddd@:::::B
                      ? dd~Y............\(lk`Y8zQx>._dddddu:::::8
                       hdddrc.........ILzcOOOOOOOO'> dddddi:::::r
                       z!dddq_.......Yi1OOOOOOOOOO?;dddddd:::::_$
                       /cdddd^h.....qiOOOOOOOOOOO>(dddddd^c::::n$
                        Q!ddddo,....?ZOOOOOOOOOOB^ddddddd,::::)v
                        ,[dddddot..I&OOOOOOOOOOBtjdddddd~}:::`*
                         b ddddd!d.U1OOOOOOOOOk>jdddddddn):::m$
                          #vddddd}qqOOOOOOOOBX~ddddddddMC:::'.
                          qcMdddddj^8zOOOOkM^!ddddddddda`:::Qq
                           qUUc<cccc+&[w:jijddddddddddnc:::$k
                           ah>IYu&>.i{!dd~f;>CBYO!ddd|L:::_o
                           Q..r[,oBY)addddY{.Yx-BQ;Md{}:::np
                           *{,>XavCjaddddddB[L}I..j(J@q):0#
                           Qd}/~+CaMdddddddd~Xk kua>..*$w<
                           od"CzOC0cddddddddu:::::'prn*.(O
                          ?id$urOCCldddddddd\::::::::k>,$)
                          f dYu">CC$ddddddd!b::::::::::)M
                          Uadd^LJ*p&UddddddX{::::::::::cB
                          Mddda-;;ad|Q dj^mZm::::::::::$(
                  $>[p{J  8@@&*qqq*_&]M||]!dx)m::::::::n
                $$OjdddkB<UxdddddddddddddLmMdrc~+Y:::::f
                +|dddddd~@;qddddddddddddddq-d#::'$W@Z::B
                odddddddd Z;dddddddddddddd_!dM:::::::::0
                fddddddddd']ddddddddddddddUddu::::::::Y*
               <cdddddddddB|!ddddddddddddj@ddu::::::::p
               M*ddddddddddbXMddddddddddd_Ydd*::::::::1
                +ddddddddddi|' dddddddd-&|dd!x::::::::: $$$
                LddddddddddviL;@c}MaBx]*-dddC}::::::::>]8.>u
                lddddddddddd^LMdv|k*q~dddddLz:::::::::o*...($
                $1ddddddddddr-dOddddddddddLZ::::::::::Uv...($
                 $LdddddddddX{:d>vddddddjOz:::::::::::\?L.ra
                  iddddddddj1:::)Ct_;xZ(b}::::::::::::1p\OU$
                  aBdddddddkp:::::::-?::::::::::::::::b\
                    !dddddd@YC@{::::::::::::::::::::::XB
                   qX~ddddnQ ?YoU+${:::::::::::::::::::aw
                    -od_+8:       B@~m)::::::::::::::::-n,
                                     qBuq:::::::::::::`}d+J
                                       q )p:::::::::::p}dj!
                                          ^[z:::::::::&ddd-
                                            !|`::::::Z-ddd@
                                             @OX::`d;Ldddd_
                                               ^LjI jddddd&
                                               dddddddddddc
                                              #dddddddddd!
                                              odddddddddd|L
                                              :dddddddddd~
                                              ~-dddddddj]r
                                              qrqdddd!nz-
                                                f)dd;b




*/
