import {AfterViewChecked, Component, OnInit, ViewChild} from "@angular/core";
import {NgForm} from "@angular/forms";
import {Scanner} from "./scanner";

@Component({
    selector: 'new-scanner-form',
    templateUrl: 'new-scanner-form.component.html',
    styleUrls: ['new-scanner-form.component.scss', '../../common.scss']
})
export class NewScannerFormComponent implements AfterViewChecked, OnInit {

    @ViewChild("newScannerForm") newScannerForm: NgForm;
    newScanner: Scanner = new Scanner(1, "x", "y", "x", false);

    ngOnInit() {

    }

    ngAfterViewChecked(): void {
    }

    getData(): Scanner {
        return this.newScanner;
    }

}
