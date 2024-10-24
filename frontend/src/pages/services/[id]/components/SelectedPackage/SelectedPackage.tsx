import SelectedPackageLayout from "./SelectedPackageLayout";
import {Package} from "@/types/Package";
import {Button, NumberFormatter} from "@mantine/core";

type SelectedPackageProps = Package;

const SelectedPackage = (props: SelectedPackageProps) => {
    const {title, price, description, deliveryDays} = props;
    const primaryTextStyle = "text-light-palette-text-primary dark:text-dark-palette-text-primary";

    return (
        <SelectedPackageLayout>
            <h2 className={`${primaryTextStyle} text-xl font-semibold mb-4`}>{title}</h2>
            <p className="mb-4 text-sm">{description}</p>
            <div className={`flex justify-between text-lg items-center ${primaryTextStyle}`}>
                <span>
                    {deliveryDays} {deliveryDays > 1 ? "Days Delivery" : "Day Delivery"}
                </span>
                <NumberFormatter
                    className={"font-bold"}
                    prefix="$"
                    value={price}
                />
            </div>
            <Button className="mt-8 !w-full !bg-cyan-500 hover:!bg-cyan-400 transition-colors duration-200">
                Order Now
            </Button>
        </SelectedPackageLayout>
    );
};

export default SelectedPackage;
