import {Injectable} from "@angular/core";
import {Scanner, ScannerMetadata} from "./scanner";
import {HttpClient} from "@angular/common/http";
import {Observable, throwError as observableThrowError} from "rxjs";

import {HTTP_JSON_OPTIONS, HTTP_GET_OPTIONS} from "@harbor/ui";
import {map, catchError, scan} from "rxjs/operators";

@Injectable()
export class ConfigScannerService {

    constructor(private http: HttpClient) {
    }

    getScanners(): Observable<Scanner[]> {
        return this.http.get('/api/scanner/registrations', HTTP_GET_OPTIONS)
            .pipe(map(response => response as Scanner[]), catchError(error => this.handleError(error)));
    }

    getScanner(id: number): Observable<Scanner> {
        return this.http.get('/api/scanner/registrations/' + id, HTTP_GET_OPTIONS)
            .pipe(map(response => response as Scanner), catchError(error => this.handleError(error)));
    }

    setAsDefault(id: number): Observable<any> {
        return this.http.put('/api/scanner/registrations/' + id + '/default', null, HTTP_JSON_OPTIONS)
            .pipe(map(() => null)
                , catchError(error => this.handleError(error)));
    }

    getDefault(): Observable<Scanner> {
        return this.http.get('/api/scanner/registrations/default', HTTP_GET_OPTIONS)
            .pipe(map(response => response as Scanner), catchError(error => this.handleError(error)));
    }

    getMetadata(id: number): Observable<ScannerMetadata> {
        return this.http.get('/api/scanner/registrations/' + id + '/metadata', HTTP_GET_OPTIONS)
            .pipe(map(response => response as ScannerMetadata), catchError(error => this.handleError(error)));
    }

    delete(id: number): Observable<any> {
        return this.http.delete('/api/scanner/registrations/' + id, HTTP_JSON_OPTIONS)
            .pipe(map(() => null)
                , catchError(error => this.handleError(error)));
    }

    create(scanner: Scanner): Observable<any> {
        return this.http.post('/api/scanner/registrations', JSON.stringify(scanner), HTTP_JSON_OPTIONS)
            .pipe(map(() => null), catchError(error => this.handleError(error)));
    }

    update(scanner: Scanner): Observable<any> {
        return this.http.put('/api/scanner/registrations/' + scanner.id, JSON.stringify(scanner), HTTP_JSON_OPTIONS)
            .pipe(map(() => null), catchError(error => this.handleError(error)));
    }

    ping(endpointURL: string): Observable<ScannerMetadata> {
        let request = JSON.stringify({"endpoint_url": endpointURL});
        return this.http.post('/api/scanner/registrations/ping', request, HTTP_JSON_OPTIONS)
            .pipe(map(response => response as ScannerMetadata), catchError(error => this.handleError(error)));
    }

    handleError(error: any): Observable<any> {
        return observableThrowError(error.error || error);
    }
}
