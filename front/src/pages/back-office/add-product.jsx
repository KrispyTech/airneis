import BackOfficeNav from "@/components/backoffice/BackOfficeNav"
import Page from "@/components/Page"
import { backoffice } from "@@/lib/constants/labels"
import { serverSideTranslations } from "next-i18next/serverSideTranslations"
import { useTranslation } from "react-i18next"
import BackOfficeForm from "@/components/backoffice/BackOfficeForm"

export const getServerSideProps = async ({ locale }) => ({ props: { ...(await serverSideTranslations(locale, ["common"])) } })
const AddProduct = () => {
  const { t } = useTranslation("common")
  const msg = "Ouiouioui"

  return (
    <Page page={backoffice}>
      <div className="flex h-full">
        <BackOfficeNav />
        <BackOfficeForm />
      </div>
    </Page>
)
}

export default AddProduct