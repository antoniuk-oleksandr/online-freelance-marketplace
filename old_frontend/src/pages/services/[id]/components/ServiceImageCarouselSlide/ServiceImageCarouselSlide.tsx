import Image from "next/image";
import {Carousel} from "@mantine/carousel";

type ServiceImageCarouselSlideProps = {
    src: string,
}

const ServiceImageCarouselSlide = (props: ServiceImageCarouselSlideProps) => {
    const {src} = props;

    return (
        <Carousel.Slide className={"aspect-video flex size-full"}>
            <Image
                priority
                sizes="full"
                src={src}
                alt="img"
                fill
                style={{ objectFit: "cover" }}
            />
        </Carousel.Slide>
    )
}

export default ServiceImageCarouselSlide;