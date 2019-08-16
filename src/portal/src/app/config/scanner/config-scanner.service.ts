import {Injectable} from "@angular/core";
import {Scanner} from "./scanner";

@Injectable()
export class ConfigScannerService {

    constructor() {}

    getScanners(): Scanner[] {
        return [];
    }

    create(s: Scanner): void {

    }
}
