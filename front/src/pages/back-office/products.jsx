import BackOfficeBoard from "@/components/BackOfficeBoard"
import BackOfficeNav from "@/components/BackOfficeNav"
import Page from "@/components/Page"
import { backoffice } from "@/constants"
import { useState } from "react"

const products = [
  {
    id: 1,
    name: "CANAPÉ BLANC",
    priority: 1,
    HTPrice: 1000,
    "in stock": "Out of stock",
    category: "Salon",
    price: 1200,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 2,
    name: "CANAPÉ SAUFA",
    priority: 1,
    HTPrice: 2020,
    "in stock": "Out of stock",
    category: "Salon",
    price: 2450,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 3,
    name: "CHAISE DES LANDES",
    priority: 1,
    HTPrice: 475,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 575,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 4,
    name: "CANAPÉ BLANC",
    priority: 1,
    HTPrice: 1000,
    "in stock": "In stock",
    category: "Salle de bain",
    price: 1200,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 5,
    name: "CANAPÉ SAUFA",
    priority: 1,
    HTPrice: 2020,
    "in stock": "In stock",
    category: "Salle de bain",
    price: 2450,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 6,
    name: "CHAISE DES LANDES",
    priority: 1,
    HTPrice: 475,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 575,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 10,
    name: "CANAPÉ BLANC",
    priority: 1,
    HTPrice: 1000,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 1200,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 20,
    name: "CANAPÉ SAUFA",
    priority: 1,
    HTPrice: 2020,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 2450,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 30,
    name: "CHAISE DES LANDES",
    priority: 1,
    HTPrice: 475,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 575,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 40,
    name: "CANAPÉ BLANC",
    priority: 1,
    HTPrice: 1000,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 1200,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 50,
    name: "CANAPÉ SAUFA",
    priority: 1,
    HTPrice: 2020,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 2450,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 60,
    name: "CHAISE DES LANDES",
    priority: 1,
    HTPrice: 475,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 575,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 14,
    name: "CANAPÉ BLANC",
    priority: 1,
    HTPrice: 1000,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 1200,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 24,
    name: "CANAPÉ SAUFA",
    priority: 1,
    HTPrice: 2020,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 2450,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 34,
    name: "CHAISE DES LANDES",
    priority: 1,
    HTPrice: 475,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 575,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 44,
    name: "CANAPÉ BLANC",
    priority: 1,
    HTPrice: 1000,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 1200,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 54,
    name: "CANAPÉ SAUFA",
    priority: 1,
    HTPrice: 2020,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 2450,
    image: "https://picsum.photos/90/70"
  },
  {
    id: 64,
    name: "CHAISE DES LANDES",
    priority: 1,
    HTPrice: 475,
    "in stock": "Out of stock",
    category: "Salle de bain",
    price: 575,
    image: "https://picsum.photos/90/70"
  }
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
    <Page page={backoffice}>
      <div className="flex h-full">
        <BackOfficeNav isOpen={isOpen} handleOpen={handleOpen} />
        <BackOfficeBoard
          items={products}
          handleOpenModal={handleOpenModal}
          handleOpen={handleOpen}
          itemId={itemId}
          isModalOpen={isModalOpen}
          handleInput={handleInput}
          handleChange={handleChange}
          selected={selected}
          columns={["image", "name", "category", "priority", "in stock", "HTPrice"]}
          pageName={"Products"}
        />
      </div>
    </Page>
  )
}

export default ProductsBoard
