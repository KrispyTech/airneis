import BackOfficeBoard from "@/components/backoffice/BackOfficeBoard"
import BackOfficeNav from "@/components/backoffice/BackOfficeNav"
import Page from "@/components/Page"
import { getCategories } from "@/utils/services/categoryServices"
import { backoffice } from "@@/lib/constants/labels"
import { useEffect, useState } from "react"

const CategoriesBoard = () => {
  const [selected, setSelected] = useState({})
  const [_, setError] = useState(false)
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [categories, setCategories] = useState([])
  const [itemId, setItemId] = useState(null)
  const handleOpenModal = event => {
    const id = event.currentTarget.getAttribute("data-id")
    setItemId(id)
    setIsModalOpen(!isModalOpen)
  }

  const handleChange = event => {
    // More things to do when filter will be done
    const id = event.currentTarget.getAttribute("data-id")

    if (id === "all") {
      setSelected({ ...selected, ...categories.reduce((xs, x) => ({ ...xs, [x.ID]: true }), {}) })

      return
    }

    setSelected({ ...selected, [id]: !selected[id] })
  }

  useEffect(() => {
    ;(async () => {
      try {
        const { data } = await getCategories()
        setCategories(data)
        setError(true)
      } catch (e) {}
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
