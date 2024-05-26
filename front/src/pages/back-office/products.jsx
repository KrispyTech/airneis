import BackOfficeBoard from "@/components/backoffice/BackOfficeBoard"
import BackOfficeNav from "@/components/backoffice/BackOfficeNav"
import Page from "@/components/Page"
import { getProducts } from "@/utils/services/productServices"
import { backoffice } from "@@/lib/constants/labels"
import { useEffect, useState } from "react"
import { serverSideTranslations } from "next-i18next/serverSideTranslations"

export const getServerSideProps = async ({ locale }) => ({ props: { ...(await serverSideTranslations(locale, ["common"])) } })

const ProductsBoard = () => {
  const { t } = useTranslation("common")
  const [selected, setSelected] = useState({})
  const [hasError, setError] = useState(false)
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

  const handleInput = () => {
    // To do to handle filter input
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

  if (hasError) {
    return <p className="bg-red-700">{t("errorMessage")}</p>
  }

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
