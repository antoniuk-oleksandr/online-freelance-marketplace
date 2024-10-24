import ServiceAboutBlockLayout from "./ServiceAboutBlockLayout";
import {Service} from "@/types/Service";

type ServiceAboutBlockProps = Service;

const ServiceAboutBlock = (props: ServiceAboutBlockProps) => {
    const {description} = props;

    return (
        <ServiceAboutBlockLayout>
            <h3 className={"text-xl font-bold text-light-palette-text-primary dark:text-dark-palette-text-primary"}>About</h3>
            <p
                className={"text-light-palette-text-secondary dark:text-dark-palette-text-secondary"}
            >{description}</p>
        </ServiceAboutBlockLayout>
    )
}

export default ServiceAboutBlock;