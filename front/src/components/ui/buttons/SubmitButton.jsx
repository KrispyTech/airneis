import Button from "@/components/ui/buttons/Button"
import { useFormikContext } from "formik"

const SubmitButton = ({ disabled, text, ...otherProps }) => {
  const { isSubmitting, isValid } = useFormikContext()

  return <Button text={text} className="rounded-lg" type="submit" variant="submit" disabled={disabled || isSubmitting || !isValid} {...otherProps} />
}

export default SubmitButton
