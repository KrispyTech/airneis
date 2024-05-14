import Page from "@/components/Page"
import SubmitButton from "@/components/ui/buttons/SubmitButton"
import FormField from "@/components/ui/forms/FormField"
import MyForm from "@/components/ui/forms/MyForm"
import useSessionStore from "@/utils/stores/sessionStore"
import routes from "@@/lib/constants/routes"
import { emailValidator, passwordValidator } from "@@/lib/validators/user"
import { Formik } from "formik"
import i18next from "i18next"
import { serverSideTranslations } from "next-i18next/serverSideTranslations"
import Link from "next/link"
import { useRouter } from "next/router"
import { useEffect, useState } from "react"
import { useTranslation } from "react-i18next"
import { object } from "yup"

export const getStaticProps = async ({ locale }) => ({ props: { ...(await serverSideTranslations(locale, ["common"])) } })

const initialValues = {
  email: "",
  password: ""
}

const validationSchema = object({
  email: emailValidator.required(),
  password: passwordValidator.required()
})

const LoginPage = () => {
  const { t } = useTranslation("common", { bindI18n: "languageChanged loaded" })
  useEffect(() => {
    i18next.reloadResources(["de", "fr"], ["ns1", "ns2"])
  }, [])

  const router = useRouter()
  const [hasError, setError] = useState(false)
  const { createUserSession } = useSessionStore()
  const handleSession = async values => {
    const { email, password } = values

    try {
      await createUserSession(email, password, router)
    } catch (e) {
      setError(true)
    }
  }

  if (hasError) {
    return <p className="bg-red-700">{t("errorMessage")}</p>
  }

  return (
    <>
      <Page className="flex justify-center">
        <div className="flex flex-col justify-center max-w-96">
          <h1 className="text-center text-3xl rounded-md">{t("auth.signIn")}</h1>
          <div className="rounded-md bg-modal m-4 border-2">
            <Formik initialValues={initialValues} validationSchema={validationSchema} onSubmit={handleSession}>
              <MyForm>
                <FormField name="email" type="email" placeholder={t("auth.placeholder.email")} label={t("auth.label.email")} />
                <FormField name="password" type="password" placeholder={t("auth.placeholder.password")} label={t("auth.label.password")} />
                <SubmitButton text={t("auth.login")} />
              </MyForm>
            </Formik>
          </div>

          <div className="flex justify-center rounded-md bg-modal m-4 border-2 p-2">
            <p className="px-1">{t("auth.newToAirneis")}</p>
            <Link href={routes.register} className="text-sky-500">
              {t("auth.createAccount")}
            </Link>
          </div>
        </div>
      </Page>
    </>
  )
}

export default LoginPage
