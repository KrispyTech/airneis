import BackOfficeModal from "@/components/BackOfficeModal"
import BackOfficeNav from "@/components/BackOfficeNav"
import BackOfficeTableHead from "@/components/BackOfficeTableHead"
import CategoryListItem from "@/components/CategoryListItem"
import Page from "@/components/Page"
import { backoffice } from "@/constants"
import routes from "@/routes"
import { Bars3Icon, PlusCircleIcon } from "@heroicons/react/24/outline"
import Link from "next/link"
import { useState } from "react"

const categories = [
  { id: 1, name: "Salon", image: "https://picsum.photos/90/70" },
  { id: 2, name: "Chambre", image: "https://picsum.photos/90/70" },
  { id: 3, name: "Salle de bain", image: "https://picsum.photos/90/70" },
  { id: 4, name: "Cuisine", image: "https://picsum.photos/90/70" },
  { id: 5, name: "Bureau", image: "https://picsum.photos/90/70" },
  { id: 6, name: "Buanderie", image: "https://picsum.photos/90/70" },
  { id: 7, name: "Jardin", image: "https://picsum.photos/90/70" },
  { id: 8, name: "Garage", image: "https://picsum.photos/90/70" }
]

const CategoriesBoard = () => {
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
      setSelected({ ...selected, ...categories.reduce((xs, x) => ({ ...xs, [x.id]: true }), {}) })

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
        <section className="flex flex-col w-full px-4 pt-3 overflow-y-scroll">
          {isModalOpen ? <BackOfficeModal handleOpenModal={handleOpenModal} id={itemId} /> : ""}
          <div className="flex justify-between">
            <h1 className="font-bold text-2xl hidden md:block">Categories</h1>
            <button onClick={handleOpen} className="md:hidden font-bold text-2xl flex items-center gap-2">
              <Bars3Icon className="h-8 w-8" strokeWidth="2" />
              Categories
            </button>
            <input placeholder="Search..." className="rounded-md drop-shadow-lg pl-2 w-40 md:w-72" onChange={handleInput} />
          </div>
          <Link href={routes.backoffice.addCategory} className="mt-4 py-0.5 px-1.5 bg-gray w-fit rounded-full border flex gap-1 items-center">
            <PlusCircleIcon className="w-6 h-6" />
            Add
          </Link>
          <table className="mt-4">
            <BackOfficeTableHead handleChange={handleChange} page="categories" />
            <tbody>
              {categories.map(category => (
                <CategoryListItem
                  key={category.id}
                  category={category}
                  selected={selected[category.id]}
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

export default CategoriesBoard
