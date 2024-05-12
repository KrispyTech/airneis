import Page from "@/components/Page"
import SubmitButton from "@/components/ui/buttons/SubmitButton"
import FormField from "@/components/ui/forms/FormField"
import MyForm from "@/components/ui/forms/MyForm"
import useSessionStore from "@/utils/stores/sessionStore"
import routes from "@@/lib/constants/routes"
import { emailValidator, passwordValidator } from "@@/lib/validators/user"
import { Formik } from "formik"
import { serverSideTranslations } from "next-i18next/serverSideTranslations"
import Link from "next/link"
import { useRouter } from "next/router"
import { object } from "yup"

export const getServerSideProps = async ({ locale }) => ({ props: { ...(await serverSideTranslations(locale, ["common"])) } })
const initialValues = {
  email: "",
  password: ""
}

const validationSchema = object({
  email: emailValidator.required(),
  password: passwordValidator.required()
})

const LoginPage = () => {
  const router = useRouter()
  const { createUserSession } = useSessionStore()
  const handleSession = async values => {
    const { email, password } = values
    try {
      await createUserSession(email, password, router)
    } catch (e) {}
  }

  return (
    <>
      <Page className="flex justify-center">
        <div className="flex flex-col justify-center max-w-96">
          <h1 className="text-center text-3xl rounded-md">Sign in to Airneis</h1>
          <div className="rounded-md bg-modal m-4 border-2">
            <Formik initialValues={initialValues} validationSchema={validationSchema} onSubmit={handleSession}>
              <MyForm>
                <FormField name="email" type="email" placeholder="Enter your e-mail" label="E-mail" />
                <FormField name="password" type="password" placeholder="Enter your password" label="Password" />
                <SubmitButton text="Login" />
              </MyForm>
            </Formik>
          </div>

          <div className="flex justify-center rounded-md bg-modal m-4 border-2 p-2">
            <p className="px-1">New to Airneis?</p>
            <Link href={routes.register} className="text-sky-500">
              Create an account
            </Link>
          </div>
        </div>
      </Page>
    </>
  )
}

export default LoginPage
