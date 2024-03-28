import { clsx } from "clsx"

const variants = {
  orange: "bg-orange-600 hover:bg-orange-700 text-black border-2 border-black py-1 px-8 uppercase"
}

const Button = props => {
  const { text, className, variant, ...otherProps } = props

  return (
    <button className={clsx(variants[variant], className)} {...otherProps}>
      {text}
    </button>
  )
}

export default Button
