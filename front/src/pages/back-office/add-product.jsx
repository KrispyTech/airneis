import BackOfficeNav from "@/components/backoffice/BackOfficeNav"
import Page from "@/components/Page"
import { useRouter } from "next/router"
import { useState, useEffect } from "react"
import { backoffice } from "@@/lib/constants/labels"
import { serverSideTranslations } from "next-i18next/serverSideTranslations"
import {
  orderOfPriorityValidator,
  priceWithoutTaxesValidator,
  priceWithTaxesValidator,
  materialsValidator,
  categoryValidator,
  descriptionValidator,
  nameValidator
} from "@@/lib/validators/product"
import { object } from "yup"
import { getCategories } from "@/utils/services/categoryServices"
import routes from "@@/lib/constants/routes"
import MyForm from "@/components/ui/forms/MyForm"
import { Formik } from "formik"
import SubmitButton from "@/components/ui/buttons/SubmitButton"
import FormField from "@/components/ui/forms/FormField"
import { useTranslation } from "react-i18next"
// The only way to have translations work on this page was to disable SSR
import NoSSR from "@/components/NoSSR"
import { createProduct } from "@/utils/services/productServices"
import FormFieldSelect from "@/components/ui/forms/FormFieldSelect"
import UploadAndDisplayImage from "@/components/ui/forms/UploadAndDisplayImages"
import AddProductForm from "@/components/ui/forms/AddProductForm"

export const getServerSideProps = async ({ locale }) => ({ props: { ...(await serverSideTranslations(locale, ["common"])) } })
const initialValues = {
  name: "",
  orderOfPriority: 0,
  priceWithoutTaxes: 0,
  priceWithTaxes: 0,
  materialsString: "",
  categoryID: 0,
  description: ""
}

const validationSchema = object({
  name: nameValidator.required(),
  orderOfPriority: orderOfPriorityValidator,
  priceWithoutTaxes: priceWithoutTaxesValidator.required(),
  priceWithTaxes: priceWithTaxesValidator.required(),
  materialsString: materialsValidator.required(),
  categoryID: categoryValidator,
  description: descriptionValidator
})

const AddProduct = () => {
  const { t } = useTranslation("common")
  const router = useRouter()
  const [hasError, setError] = useState(false)
  const [categories, setCategories] = useState([])
  const [images, setImages] = useState(["/no_img.jpeg"])
  const handleProductCreation = async values => {
    const { name, orderOfPriority, priceWithoutTaxes, priceWithTaxes, materialsString, categoryID, description } = values
    const materials = materialsString.split(",").map(material => ({ name: material }))
    const category = categories.filter(categorie => categorie.ID === categoryID)[0]
    const slug = name.split(" ").join("_").toLowerCase()

    try {
      const product = {
        name,
        slug,
        categoryID: parseInt(categoryID),
        orderOfPriority: parseInt(orderOfPriority),
        priceWithoutTaxes: parseInt(priceWithoutTaxes),
        priceWithTaxes: parseInt(priceWithTaxes),
        materials,
        category,
        description,
        thumbnailURL: images[0]
      }
      await createProduct(product)
      router.push(routes.backoffice.products)
    } catch (e) {
      setError(true)
    }
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
      <NoSSR>
        <div className="flex h-full">
          <BackOfficeNav />
          <section className="flex flex-col w-full px-4 pt-3 overflow-y-scroll">
            <Formik initialValues={initialValues} validationSchema={validationSchema} onSubmit={handleProductCreation}>
              <AddProductForm images={images} setImages={setImages} categories={categories} t={t}/>
            </Formik>
          </section>
        </div>
      </NoSSR>
    </Page>
  )
}

export default AddProduct
