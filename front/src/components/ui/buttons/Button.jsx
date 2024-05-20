import { clsx } from "clsx"

const variants = {
  orange: "bg-button-normal hover:bg-button-hover text-black border-4 border-black py-1 px-8 uppercase disabled:opacity-50 disabled:border-slate-600",
  submit: "bg-button-normal hover:bg-button-hover rounded-md bold-xl text-black border-2 border-black py-2 px-8 uppercase"
}

const Button = ({ text, className, variant, type, ...otherProps }) => (
  <button type={type} className={clsx(variants[variant], className)} {...otherProps}>
    {text}
  </button>
)

export default Button
