import ServiceBreadcrumbsLayout from "./ServiceBreadcrumbsLayout";
import Link from "next/link";
import {Breadcrumbs} from "@mantine/core";
import {Service} from "@/types/Service";

type ServiceBreadcrumbsProps = Service;

const ServiceBreadcrumbs = (props: ServiceBreadcrumbsProps) => {
    const {category} = props;

    const items = [
        {title: "Home", href: "/"},
        {title: category.name, href: `/search?category=${category.id}`},
    ].map((item) => (
        <Link
            className={"hover:underline ease-out duration-200"}
            href={item.href} key={item.title}
        >
            {item.title}
        </Link>
    ));

    return (
        <ServiceBreadcrumbsLayout>
            <Breadcrumbs>{items}</Breadcrumbs>
        </ServiceBreadcrumbsLayout>
    )
}

export default ServiceBreadcrumbs;