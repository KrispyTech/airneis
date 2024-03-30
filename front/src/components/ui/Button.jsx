import { clsx } from "clsx"

const variants = {
  orange: "bg-orange-600 hover:bg-orange-700 text-black border-4 border-black py-1 px-8 uppercase disabled:opacity-50 disabled:border-slate-600"
}

const Button = ({ text, className, variant, ...otherProps }) => (
  <button className={clsx(variants[variant], className)} {...otherProps}>
    {text}
  </button>
)

export default Button
