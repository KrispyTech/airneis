import Header from "@/components/Header"
import { clsx } from "clsx"

const Page = ({ children, className, ...otherProps }) => (
  <>
    <Header />
    <main className={clsx("bg-primary", className)} {...otherProps}>
      {children}
    </main>
  </>
)

export default Page
