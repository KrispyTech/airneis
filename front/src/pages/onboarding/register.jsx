import { serverSideTranslations } from "next-i18next/serverSideTranslations"
import routes from "@@/lib/constants/routes"
import Page from "@/components/Page"
import { object } from "yup"
import { useRouter } from "next/router"
import { Formik } from "formik"
import MyForm from "@/components/ui/forms/MyForm"
import { emailValidator, firstNameValidator, lastNameValidator, passwordValidator } from "@@/lib/validators/user"
import FormField from "@/components/ui/forms/FormField"
import SubmitButton from "@/components/ui/buttons/SubmitButton"

const initialValues = {
    firstName: "",
    lastName: "",
    email: "",
    password: ""
}

const validationSchema = object({
  firstName: firstNameValidator.required(),
  lastName: lastNameValidator.required(),
  email: emailValidator.required(),
  password: passwordValidator.required()
})

export const getServerSideProps = async ({ locale }) => ({ props: { ...(await serverSideTranslations(locale, ["common"])) } })
const RegisterPage = () => {
  const router = useRouter()
  const handleRegister = () => {
    router.push(routes.home)
  }

  return (
    <Page>
      <div className="w-full h-1/4 sm:h-1/3 bg-[url('https://picsum.photos/1800/350')] bg-center flex items-center justify-center">
        <Formik initialValues={initialValues} validationSchema={validationSchema} onSubmit={handleRegister}>
          <MyForm>
            <FormField name="firstName" type="text" placeholder="Enter your first name" label="First Name" />
            <FormField name="lastName" type="text" placeholder="Enter your last name" label="Last Name" />
            <FormField name="email" type="email" placeholder="Enter your e-mail" label="E-mail" />
            <FormField name="password" type="password" placeholder="Enter your password" label="Password" />
            <SubmitButton>"Register"</SubmitButton>
          </MyForm>
        </Formik>
      </div>
    </Page>
  )
}

export default RegisterPage
