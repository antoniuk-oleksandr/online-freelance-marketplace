import Image from "next/image";
import {Carousel} from "@mantine/carousel";

type ServiceImageCarouselSlideProps = {
    src: string,
}

const ServiceImageCarouselSlide = (props: ServiceImageCarouselSlideProps) => {
    const {src} = props;

    return (
        <Carousel.Slide className={"aspect-video"}>
            <Image
                width={1920}
                height={1080}
                alt={"img"}
                src={src}
            />
        </Carousel.Slide>
    )
}

export default ServiceImageCarouselSlide;