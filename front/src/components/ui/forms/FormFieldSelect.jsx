import Alert from "@/components/Alert"
import { Field } from "formik"
import clsx from "clsx"
import { useField } from "formik"

const FormFieldSelect = props => {
  const { as: Component = Field, name, label, options, children, className, ...otherProps } = props
  const [field] = useField(name)

  return (
  <label className={clsx("flex flex-row gap-2", className)}>
    {label && <span className="content-center font-semibold">{label}</span>}
    <Component name={name} component="select" {...field} {...otherProps} className="border-2 p-2 bg-white flex-row">
      {options.map(option => (
        <option key={option.ID} value={option.ID}>{option.name}</option>
        ))}
    </Component>
    {children && <Alert>{children}</Alert>}
  </label>
  )
}

export default FormFieldSelect