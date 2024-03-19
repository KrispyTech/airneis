import BackOfficeBoard from "@/components/BackOfficeBoard"
import BackOfficeNav from "@/components/BackOfficeNav"
import Page from "@/components/Page"
import { backoffice } from "@/constants"
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
        <BackOfficeNav isOpen={isOpen} handleOpen={handleOpen} columns={["image", "name"]} />
        <BackOfficeBoard
          items={categories}
          handleOpenModal={handleOpenModal}
          handleOpen={handleOpen}
          itemId={itemId}
          isModalOpen={isModalOpen}
          handleInput={handleInput}
          handleChange={handleChange}
          selected={selected}
          columns={["image", "name"]}
          pageName={"Categories"}
        />
      </div>
    </Page>
  )
}

export default CategoriesBoard
