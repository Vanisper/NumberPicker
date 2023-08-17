export interface IConfig {
    Category: ICategory;
}

interface ICategory {
    [x: string]: number[];
}