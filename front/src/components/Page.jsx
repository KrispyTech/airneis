import Header from "@/components/Header"
import { clsx } from "clsx"

const Page = ({ children, className, page, ...otherProps }) => (
  <div className="flex flex-col h-screen">
    <Header page={page} />
    <main className={clsx("bg-primary grow overflow-y-auto", className)} {...otherProps}>
      {children}
    </main>
  </div>
)

export default Page
