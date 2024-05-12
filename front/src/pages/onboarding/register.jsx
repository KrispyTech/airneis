import Page from "@/components/Page"
import SubmitButton from "@/components/ui/buttons/SubmitButton"
import FormField from "@/components/ui/forms/FormField"
import MyForm from "@/components/ui/forms/MyForm"
import { createUser } from "@/utils/services/userServices"
import routes from "@@/lib/constants/routes"
import { emailValidator, firstNameValidator, lastNameValidator, passwordValidator } from "@@/lib/validators/user"
import { Formik } from "formik"
import { serverSideTranslations } from "next-i18next/serverSideTranslations"
import Link from "next/link"
import { useRouter } from "next/router"
import { useState } from "react"
import { useTranslation } from "react-i18next"
import { object } from "yup"

export const getServerSideProps = async ({ locale }) => ({ props: { ...(await serverSideTranslations(locale, ["common"])) } })
const initialValues = {
  firstName: "",
  lastName: "",
  email: "",
  password: ""
}

const validationSchema = object({
  email: emailValidator.required(),
  password: passwordValidator.required(),
  firstName: firstNameValidator.required(),
  lastName: lastNameValidator.required()
})

const RegisterPage = () => {
  const { t } = useTranslation("common")
  const router = useRouter()
  const [hasError, setError] = useState(false)
  const handleUserCreation = async values => {
    const { firstName, lastName, email, password } = values

    try {
      const user = { firstName, lastName, email, password }
      await createUser(user)
      router.push(routes.home)
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
          <h1 className="text-center text-3xl rounded-md">Create your account! Welcome to Airneis 😄</h1>
          <div className="rounded-md bg-modal m-4 border-2">
            <Formik initialValues={initialValues} validationSchema={validationSchema} onSubmit={handleUserCreation}>
              <MyForm>
                <FormField name="firstName" type="text" placeholder="Enter your first name" label="First Name" />
                <FormField name="lastName" type="text" placeholder="Enter your last name" label="Last Name" />
                <FormField name="email" type="email" placeholder="Enter your e-mail" label="E-mail" />
                <FormField name="password" type="password" placeholder="Enter your password" label="Password" />
                <SubmitButton text="Create account" />
              </MyForm>
            </Formik>
          </div>
          <div className="flex justify-center rounded-md bg-modal m-4 border-2 p-2">
            <p className="px-1">Have an account?</p>
            <Link href={routes.login} className="text-sky-500">
              Login
            </Link>
          </div>
        </div>
      </Page>
    </>
  )
}

export default RegisterPage
