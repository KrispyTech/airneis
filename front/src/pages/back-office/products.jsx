import BackOfficeModal from "@/components/BackOfficeModal"
import BackOfficeNav from "@/components/BackOfficeNav"
import BackOfficeTableHead from "@/components/BackOfficeTableHead"
import Page from "@/components/Page"
import ProductListItem from "@/components/ProductListItem"
import { Bars3Icon, PlusCircleIcon } from "@heroicons/react/24/outline"
import Link from "next/link"
import { useState } from "react"

const products = [
  { id: 1, name: "CANAPÉ BLANC", price: 1200, image: "https://picsum.photos/90/70" },
  { id: 2, name: "CANAPÉ SAUFA", price: 2450, image: "https://picsum.photos/90/70" },
  { id: 3, name: "CHAISE DES LANDES", price: 575, image: "https://picsum.photos/90/70" },
  { id: 4, name: "CANAPÉ BLANC", price: 1200, image: "https://picsum.photos/90/70" },
  { id: 5, name: "CANAPÉ SAUFA", price: 2450, image: "https://picsum.photos/90/70" },
  { id: 6, name: "CHAISE DES LANDES", price: 575, image: "https://picsum.photos/90/70" },
  { id: 10, name: "CANAPÉ BLANC", price: 1200, image: "https://picsum.photos/90/70" },
  { id: 20, name: "CANAPÉ SAUFA", price: 2450, image: "https://picsum.photos/90/70" },
  { id: 30, name: "CHAISE DES LANDES", price: 575, image: "https://picsum.photos/90/70" },
  { id: 40, name: "CANAPÉ BLANC", price: 1200, image: "https://picsum.photos/90/70" },
  { id: 50, name: "CANAPÉ SAUFA", price: 2450, image: "https://picsum.photos/90/70" },
  { id: 60, name: "CHAISE DES LANDES", price: 575, image: "https://picsum.photos/90/70" },
  { id: 14, name: "CANAPÉ BLANC", price: 1200, image: "https://picsum.photos/90/70" },
  { id: 24, name: "CANAPÉ SAUFA", price: 2450, image: "https://picsum.photos/90/70" },
  { id: 34, name: "CHAISE DES LANDES", price: 575, image: "https://picsum.photos/90/70" },
  { id: 44, name: "CANAPÉ BLANC", price: 1200, image: "https://picsum.photos/90/70" },
  { id: 54, name: "CANAPÉ SAUFA", price: 2450, image: "https://picsum.photos/90/70" },
  { id: 64, name: "CHAISE DES LANDES", price: 575, image: "https://picsum.photos/90/70" }
]

const ProductsBoard = () => {
  const [isOpen, setIsOpen] = useState(false)
  const handleOpen = () => setIsOpen(!isOpen)
  const [selected, setSelected] = useState({})
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [itemId, setItemId] = useState(null)
  const handleOpenModal = event => {
    const id = event.currentTarget.getAttribute("data-id")
    setItemId(id)
    setIsModalOpen(!isModalOpen)
  }

  const handleChange = event => {
    // More things to do when filter will be done.
    const id = event.currentTarget.getAttribute("data-id")

    if (id === "all") {
      setSelected({ ...selected, ...products.reduce((xs, x) => ({ ...xs, [x.id]: true }), {}) })

      return
    }

    setSelected({ ...selected, [id]: !selected[id] })
  }

  const handleInput = () => {
    // To do
  }

  return (
    <Page page="backoffice">
      <div className="flex h-full">
        <BackOfficeNav isOpen={isOpen} handleOpen={handleOpen} />
        <section className="flex flex-col w-full px-4 pt-3 overflow-y-scroll">
          {isModalOpen ? <BackOfficeModal handleOpenModal={handleOpenModal} id={itemId} /> : ""}
          <div className="flex justify-between">
            <h1 className="font-bold text-2xl hidden md:block">Products</h1>
            <button onClick={handleOpen} className="md:hidden font-bold text-2xl flex items-center gap-2">
              <Bars3Icon className="h-8 w-8" strokeWidth="2" />
              Products
            </button>
            <input placeholder="Search..." className="rounded-md drop-shadow-lg pl-2 w-40 md:w-72" onChange={handleInput} />
          </div>
          <Link href="/backoffice/add-product" className="mt-4 py-0.5 px-1.5 bg-gray w-fit rounded-full border flex gap-1 items-center">
            <PlusCircleIcon className="w-6 h-6" />
            Add
          </Link>
          <table className="mt-4">
            <BackOfficeTableHead handleChange={handleChange} />
            <tbody>
              {products.map(product => (
                <ProductListItem
                  key={product.id}
                  product={product}
                  selected={selected[product.id]}
                  handleOpenModal={handleOpenModal}
                  handleChange={handleChange}
                />
              ))}
            </tbody>
          </table>
        </section>
      </div>
    </Page>
  )
}

export default ProductsBoard
