import BackOfficeBoard from "@/components/backoffice/BackOfficeBoard"
import BackOfficeNav from "@/components/backoffice/BackOfficeNav"
import Page from "@/components/Page"
import { getCategories } from "@/utils/services/categoryServices"
import { backoffice } from "@@/lib/constants/labels"
import { useEffect, useState } from "react"
import { serverSideTranslations } from "next-i18next/serverSideTranslations"
import { useTranslation } from "react-i18next"

export const getServerSideProps = async ({ locale }) => ({ props: { ...(await serverSideTranslations(locale, ["common"])) } })
const CategoriesBoard = () => {
  const { t } = useTranslation("common")
  const [selected, setSelected] = useState({})
  const [hasError, setError] = useState(false)
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
      setSelected({ ...selected, ...categories.reduce((xs, x) => ({ ...xs, [x.ID]: true }), {}) })

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
        setCategories(data)
      } catch (e) {
        setError(true)
      }
    })()
  }, [setCategories])

  if (hasError) {
    return <p className="bg-red-700">{t("errorMessage")}</p>
  }

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
