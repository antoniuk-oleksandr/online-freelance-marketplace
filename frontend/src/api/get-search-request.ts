import {getHost} from "@/utils/utils.ts";

export const getSearchRequest = () => {
    const services = [
        {
            "id": 1,
            "created_at": "2024-11-02T21:20:45.403799Z",
            "description": "Professional web development services tailored to your needs.",
            "title": "Custom Website Development",
            "category_id": 1,
            "freelancer_id": 1,
            "image": "http://localhost:8030/files/image_1.jpg",
            "reviewsCount": 1,
            "rating": 1,
            "minPrice": 49.99
        },
        {
            "id": 2,
            "created_at": "2024-11-02T21:20:45.403799Z",
            "description": "Modern, mobile-friendly web designs.",
            "title": "Responsive Web Design",
            "category_id": 1,
            "freelancer_id": 1,
            "image": "",
            "reviewsCount": 0,
            "rating": 0,
            "minPrice": 99.99
        },
        {
            "id": 3,
            "created_at": "2024-11-02T21:20:45.403799Z",
            "description": "Build a powerful online store with our e-commerce solutions.",
            "title": "E-commerce Web Development",
            "category_id": 1,
            "freelancer_id": 1,
            "image": "",
            "reviewsCount": 1,
            "rating": 5,
            "minPrice": 29.99
        },
        {
            "id": 4,
            "created_at": "2024-11-02T21:20:45.403799Z",
            "description": "Unique logo design services to represent your brand.",
            "title": "Logo Design",
            "category_id": 2,
            "freelancer_id": 1,
            "image": "",
            "reviewsCount": 1,
            "rating": 2,
            "minPrice": 99.99
        },
        {
            "id": 5,
            "created_at": "2024-11-02T21:20:45.403799Z",
            "description": "Complete brand identity design for businesses.",
            "title": "Brand Identity Design",
            "category_id": 2,
            "freelancer_id": 1,
            "image": "",
            "reviewsCount": 0,
            "rating": 0,
            "minPrice": 59.99
        },
        {
            "id": 6,
            "created_at": "2024-11-02T21:20:45.403799Z",
            "description": "Professional designs for brochures, flyers, and more.",
            "title": "Graphic Design for Print",
            "category_id": 2,
            "freelancer_id": 1,
            "image": "",
            "reviewsCount": 1,
            "rating": 4,
            "minPrice": 49.99
        },
        {
            "id": 7,
            "created_at": "2024-11-02T21:20:45.403799Z",
            "description": "Get your website to rank higher on search engines.",
            "title": "SEO Optimization",
            "category_id": 4,
            "freelancer_id": 1,
            "image": "",
            "reviewsCount": 0,
            "rating": 0,
            "minPrice": 149.99
        }
    ]

    return {
        data: {
            services: [...services,...services,...services,...services,...services,...services]
        },
        status: 200
    }
}