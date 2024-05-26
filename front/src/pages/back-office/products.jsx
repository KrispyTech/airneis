import BackOfficeBoard from "@/components/backoffice/BackOfficeBoard"
import BackOfficeNav from "@/components/backoffice/BackOfficeNav"
import Page from "@/components/Page"
import { getProducts } from "@/utils/services/productServices"
import { backoffice } from "@@/lib/constants/labels"
import { useEffect, useState } from "react"

const ProductsBoard = () => {
  const [selected, setSelected] = useState({})
  const [_, setError] = useState(false)
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [products, setProducts] = useState([])
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

  useEffect(() => {
    ;(async () => {
      try {
        const { data } = await getProducts()
        setProducts(data)
      } catch (e) {
        setError(true)
      }
    })()
  }, [setProducts])

  return (
    <Page page={backoffice}>
      <div className="flex h-full">
        <BackOfficeNav />
        <BackOfficeBoard
          items={products}
          handleOpenModal={handleOpenModal}
          itemId={itemId}
          isModalOpen={isModalOpen}
          handleInput={handleInput}
          handleChange={handleChange}
          selected={selected}
          columns={["name", "categoryID", "orderOfPriority", "outOfStock", "priceWithoutTaxes"]}
          pageName={"Products"}
        />
      </div>
    </Page>
  )
}

export default ProductsBoard
