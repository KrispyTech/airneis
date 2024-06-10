import MyForm from "@/components/ui/forms/MyForm"
import { Formik } from "formik"
import SubmitButton from "@/components/ui/buttons/SubmitButton"
import FormField from "@/components/ui/forms/FormField"

const BackOfficeForm = () => (
  <section className="flex flex-col w-full px-4 pt-3 overflow-y-scroll">
    <Formik>
      <MyForm>
        <FormField name="email" type="email" placeholder={"yo"} label={"yoyo"} />
        <FormField name="password" type="password" placeholder={"yo"} label={"yo"} />
        <SubmitButton text={"validate"} />
      </MyForm>
    </Formik>
  </section>
)

export default BackOfficeForm