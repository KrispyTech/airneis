import Page from "@/components/Page"
import Button from "@/components/ui/Button"
import { TrashIcon } from "@heroicons/react/24/outline"
import Image from "next/image"
import { useEffect, useState } from "react"

const cart = [
  { product: { id: 2, name: "Canapé blanc", price: 1200 }, quantity: 1 },
  { product: { id: 1, name: "Canapé sofa", price: 2450 }, quantity: 2 },
  { product: { id: 3, name: "Chaise des landes", price: 575 }, quantity: 1 }
]

// export const getServerSideProps = async ({ locale }) => ({ props: { cart, ...(await serverSideTranslations(locale, ["common"])) } })
const Cart = () => {
  // const { t } = useTranslation()
  const [total, setTotal] = useState(0)
  const [TVA, setTVA] = useState(0)

  useEffect(() => {
    if (cart.length >= 1) {
      const Total = cart.reduce((xs, x) => xs + x.product.price * x.quantity, 0)
      const tva = (Total * 20) / 100
      setTotal(Total + tva)
      setTVA(tva)
    }
  }, [setTotal, setTVA])

  return (
    <Page>
      <div className="flex flex-col w-full items-center">
        <h1 className="font-bold text-4xl my-6">Cart</h1>
        <div className="flex flex-col items-center md:flex-row md:justify-evenly md:items-start w-full">
          <ul className="flex flex-col gap-6 md:gap-14">
            {cart.length >= 1 ? (
              cart.map(item => (
                <li key={item.id} className="flex gap-6">
                  <div className="rounded-md border-2 border-brown">
                    <Image
                      src="https://picsum.photos/154/150"
                      alt="product image"
                      width={150}
                      height={150}
                      className="rounded-[0.24rem] drop-shadow-box"
                    />
                  </div>
                  <div className="flex flex-col gap-2.5 my-auto">
                    <h2 className="font-semibold uppercase">{item.product.name}</h2>
                    <p>{item.product.price}€</p>
                    <input placeholder="Qty" className="w-10 font-bold pl-3" value={item.quantity} />
                    <TrashIcon className="w-7 h-7 text-red-400" />
                  </div>
                </li>
              ))
            ) : (
              <li>No products in cart</li>
            )}
          </ul>
          <div className="flex flex-col w-2/3 sm:w-1/2 md:w-1/3 lg:w-1/4 mt-6 gap-2">
            <p className="flex justify-between font-bold text-2xl uppercase">
              Total<span>{total}€</span>
            </p>
            <p className="flex justify-between font-semibold text-lg text-darkGray">
              TVA<span>{TVA}€</span>
            </p>
            <Button text="Make order" variant="orange" className="mt-8 py-1.5 font-semibold text-xl" />
          </div>
        </div>
      </div>
    </Page>
  )
}

export default Cart
