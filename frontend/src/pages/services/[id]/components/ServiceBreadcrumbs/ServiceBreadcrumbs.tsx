import ServiceBreadcrumbsLayout from "./ServiceBreadcrumbsLayout";
import Link from "next/link";
import {Breadcrumbs} from "@mantine/core";
import {Service} from "@/types/Service";

type ServiceBreadcrumbsProps = Service;

const ServiceBreadcrumbs = (props: ServiceBreadcrumbsProps) => {
    const {category, title} = props;

    const items = [
        {title: "Home", href: "/"},
        {title: title, href: `/search?category=${category}`},
    ].map((item) => (
        <Link href={item.href} key={item.title}>
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