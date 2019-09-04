import {Component, ViewChild} from "@angular/core";
import {NgForm} from "@angular/forms";
import {Scanner} from "./scanner";

@Component({
    selector: 'new-scanner-form',
    templateUrl: 'new-scanner-form.component.html',
    styleUrls: ['new-scanner-form.component.scss', '../../common.scss']
})
export class NewScannerFormComponent {

    @ViewChild("newScannerForm") newScannerForm: NgForm;
    newScanner: Scanner = new Scanner();

    getData(): Scanner {
        return this.newScanner;
    }

    setData(scanner: Scanner): void {
        this.newScanner = scanner;
    }

}
