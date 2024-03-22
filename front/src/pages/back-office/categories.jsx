import BackOfficeBoard from "@/components/BackOfficeBoard"
import BackOfficeNav from "@/components/BackOfficeNav"
import Page from "@/components/Page"
import { backoffice } from "@/constants"
import { getCategories } from "@/utils/services/CategoryServices"
import { useEffect, useState } from "react"

const CategoriesBoard = () => {
  const [selected, setSelected] = useState({})
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [categories, setCategories] = useState([])
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

  useEffect(() => {
    ;(async () => {
      try {
        const { data } = await getCategories()
        console.log("Result of get categories", data)
      } catch (e) {
        console.log(e)
      }
    })()
  }, [setCategories])

  return (
    <Page page={backoffice}>
      <div className="flex h-full">
        <BackOfficeNav />
        <BackOfficeBoard
          items={categories}
          handleOpenModal={handleOpenModal}
          itemId={itemId}
          isModalOpen={isModalOpen}
          handleInput={handleInput}
          handleChange={handleChange}
          selected={selected}
          columns={["name"]}
          pageName={"Categories"}
        />
      </div>
    </Page>
  )
}

export default CategoriesBoard
