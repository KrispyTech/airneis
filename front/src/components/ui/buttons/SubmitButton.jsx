import Button from "@/web/ui/components/Button"
import { useFormikContext } from "formik"

const SubmitButton = ({ disabled, ...otherProps }) => {
  const { isSubmitting, isValid } = useFormikContext()

  return <Button type="submit" variant="submit" disabled={disabled || isSubmitting || !isValid} {...otherProps} />
}

export default SubmitButton
