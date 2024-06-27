import BackOfficeNav from "@/components/backoffice/BackOfficeNav"
import Page from "@/components/Page"
import { useRouter } from "next/router"
import { useState, useEffect } from "react"
import { backoffice } from "@@/lib/constants/labels"
import { serverSideTranslations } from "next-i18next/serverSideTranslations"
import { orderOfPriorityValidator, priceWithoutTaxesValidator, priceWithTaxesValidator, materialsValidator, categoryValidator, descriptionValidator, nameValidator } from "@@/lib/validators/product"
import { object } from "yup"
import { getCategories } from "@/utils/services/categoryServices"
import routes from "@@/lib/constants/routes"
import MyForm from "@/components/ui/forms/MyForm"
import { Formik, Field } from "formik"
import SubmitButton from "@/components/ui/buttons/SubmitButton"
import FormField from "@/components/ui/forms/FormField"
import { useTranslation } from "react-i18next"
// The only way to have translations work on this page was to disable SSR
import NoSSR from "@/components/NoSSR"
import { createProduct } from "@/utils/services/productServices"
import clsx from "clsx"
import Input from "@/components/ui/Input"
import UploadAndDisplayImage from "@/components/ui/forms/UploadImages"

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
  description: descriptionValidator,
})

const AddProduct = () => {
  const { t } = useTranslation("common")
  const router = useRouter()
  const [hasError, setError] = useState(false)
  const [categories, setCategories] = useState([])
  const handleProductCreation = async values => {
    const { name, orderOfPriority, priceWithoutTaxes, priceWithTaxes, materialsString, categoryID, description } = values
    const materials = materialsString.split(',').map((material) => ({"name": material}))
    const category = categories.filter((categorie) => categorie.ID == categoryID)[0]
    const slug = name.split(' ').join('_').toLowerCase()
    console.log(typeof categoryID)

    try {
      const product = { name, slug, categoryID: parseInt(categoryID), orderOfPriority: parseInt(categoryID), priceWithoutTaxes: parseInt(priceWithoutTaxes), priceWithTaxes: parseInt(priceWithTaxes), materials, category, description }
      console.log(product)
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
              <MyForm>
                < div className="flex justify-between gap-x-16">
                  <div className="flex flex-col items-center justify-center">
                    <UploadAndDisplayImage/>
                  </div>
                  <div className="flex flex-1 flex-col justify-start">
                  <FormField name="name" type="text" placeholder={t("backoffice.placeholder.name")} label={t("backoffice.label.name")} />
                  <FormField name="orderOfPriority" type="number" placeholder={t("backoffice.placeholder.orderOfPriority")} label={t("backoffice.label.orderOfPriority")} />
                  <FormField name="priceWithoutTaxes" type="number" placeholder={t("backoffice.placeholder.priceWithoutTaxes")} label={t("backoffice.label.priceWithoutTaxes")} />
                  <FormField name="priceWithTaxes" type="number" placeholder={t("backoffice.placeholder.priceWithTaxes")} label={t("backoffice.label.priceWithTaxes")} />
                  <FormField name="materialsString" type="text" placeholder={t("backoffice.placeholder.materials")} label={t("backoffice.label.materials")} />
                  <label className={clsx("flex flex-col gap-2")}>
                  {t("backoffice.label.category") && <span className="font-semibold">{t("backoffice.label.category")}</span>}
                  <Field name="categoryID" component="select" placeholder={t("backoffice.placeholder.category")} label={t("backoffice.label.category")} className="border-2 p-2 bg-white">
                    {categories.map(categorie => (
                      <option key={categorie.ID} value={categorie.ID}>{categorie.name}</option>
                      ))}
                  </Field>
                  </label>
                  </div>
                </div>
                <FormField name="description" type="textarea" placeholder={t("backoffice.placeholder.description")} label={t("backoffice.label.description")} />
                <SubmitButton text={"validate"} />
              </MyForm>
            </Formik>
          </section>
        </div>
      </NoSSR>
    </Page>
)
}

export default AddProduct