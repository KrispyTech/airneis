import BackOfficeNav from "@/components/backoffice/BackOfficeNav"
import Page from "@/components/Page"
import { useRouter } from "next/router"
import { useState, useEffect } from "react"
import { backoffice } from "@@/lib/constants/labels"
import { serverSideTranslations } from "next-i18next/serverSideTranslations"
import { priorityValidator, priceETValidator, priceITValidator, materialsValidator, categoryValidator, descriptionValidator, nameValidator } from "@@/lib/validators/product"
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

export const getServerSideProps = async ({ locale }) => ({ props: { ...(await serverSideTranslations(locale, ["common"])) } })
const initialValues = {
  name: "",
  priority: 0,
  priceET: 0,
  priceIT: 0,
  materialsString: "",
  categoryID: 0,
  description: ""
}

const validationSchema = object({
  name: nameValidator.required(),
  priority: priorityValidator,
  priceET: priceETValidator.required(),
  priceIT: priceITValidator.required(),
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
    const { name, priority, priceET, priceIT, materialsString, categoryID, description } = values
    const materials = materialsString.split(',').map((material) => ({"name": material}))
    const category = categories.filter((categorie) => categorie.ID == categoryID)[0]

    try {
      const product = { name, priority, priceET, priceIT, materials, category, description }
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
                <FormField name="name" type="text" placeholder={t("backoffice.placeholder.name")} label={t("backoffice.label.name")} />
                <FormField name="priority" type="number" placeholder={t("backoffice.placeholder.priority")} label={t("backoffice.label.priority")} />
                <FormField name="priceET" type="number" placeholder={t("backoffice.placeholder.priceET")} label={t("backoffice.label.priceET")} />
                <FormField name="priceIT" type="number" placeholder={t("backoffice.placeholder.priceIT")} label={t("backoffice.label.priceIT")} />
                <FormField name="materialsString" type="text" placeholder={t("backoffice.placeholder.materials")} label={t("backoffice.label.materials")} />
                <label className={clsx("flex flex-col gap-2")}>
                {t("backoffice.label.category") && <span className="font-semibold">{t("backoffice.label.category")}</span>}
                <Field name="categoryID" component="select" placeholder={t("backoffice.placeholder.category")} label={t("backoffice.label.category")} >
                  {categories.map(categorie => (
                    <option key={categorie.ID} value={categorie.ID}>{categorie.ID}</option>
                    ))}
                </Field>
                </label>
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