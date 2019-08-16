import {Component, ViewChild, OnInit} from "@angular/core";
import {Scanner} from "./scanner";
import {NewScannerModalComponent} from "./new-scanner-modal.component";
import {ConfigScannerService} from "./config-scanner.service";

@Component({
    selector: 'config-scanner',
    templateUrl: "config-scanner.component.html",
    styleUrls: ['./config-scanner.component.scss', '../config.component.scss']
})
export class ConfigurationScannerComponent implements OnInit {

    scanners: Scanner[] = [];
    selectedRow: Scanner[] = [];

    @ViewChild(NewScannerModalComponent)
    newScannerDialog: NewScannerModalComponent;

    constructor(
        private configScannerService: ConfigScannerService
    ) {
    }

    ngOnInit() {
        let scanners = this.configScannerService.getScanners();
        console.log(scanners);

        this.scanners.push(new Scanner(1, 'Clair', 'CoreOS', 'https://harbor-scanner-clair:8080', true));
        this.scanners.push(new Scanner(2, 'Microscanner', 'Aqua Security', 'https://harbor-scanner-microscanner:8080', false));
        this.scanners.push(new Scanner(2, 'Anchore Enterprise', 'Anchore', 'https://harbor-scanner-anchore:8080', false));
    }

    addNewScanner(): void {
        this.newScannerDialog.open();
    }

}
