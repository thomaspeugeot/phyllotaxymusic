// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { BezierAPI } from './bezier-api'
import { Bezier, CopyBezierToBezierAPI } from './bezier'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { ShapeCategoryAPI } from './shapecategory-api'

@Injectable({
  providedIn: 'root'
})
export class BezierService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  BezierServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private beziersUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.beziersUrl = origin + '/api/github.com/thomaspeugeot/phyllotaxymusic/go/v1/beziers';
  }

  /** GET beziers from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<BezierAPI[]> {
    return this.getBeziers(GONG__StackPath, frontRepo)
  }
  getBeziers(GONG__StackPath: string, frontRepo: FrontRepo): Observable<BezierAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<BezierAPI[]>(this.beziersUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<BezierAPI[]>('getBeziers', []))
      );
  }

  /** GET bezier by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<BezierAPI> {
    return this.getBezier(id, GONG__StackPath, frontRepo)
  }
  getBezier(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<BezierAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.beziersUrl}/${id}`;
    return this.http.get<BezierAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched bezier id=${id}`)),
      catchError(this.handleError<BezierAPI>(`getBezier id=${id}`))
    );
  }

  // postFront copy bezier to a version with encoded pointers and post to the back
  postFront(bezier: Bezier, GONG__StackPath: string): Observable<BezierAPI> {
    let bezierAPI = new BezierAPI
    CopyBezierToBezierAPI(bezier, bezierAPI)
    const id = typeof bezierAPI === 'number' ? bezierAPI : bezierAPI.ID
    const url = `${this.beziersUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<BezierAPI>(url, bezierAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<BezierAPI>('postBezier'))
    );
  }
  
  /** POST: add a new bezier to the server */
  post(bezierdb: BezierAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<BezierAPI> {
    return this.postBezier(bezierdb, GONG__StackPath, frontRepo)
  }
  postBezier(bezierdb: BezierAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<BezierAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<BezierAPI>(this.beziersUrl, bezierdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted bezierdb id=${bezierdb.ID}`)
      }),
      catchError(this.handleError<BezierAPI>('postBezier'))
    );
  }

  /** DELETE: delete the bezierdb from the server */
  delete(bezierdb: BezierAPI | number, GONG__StackPath: string): Observable<BezierAPI> {
    return this.deleteBezier(bezierdb, GONG__StackPath)
  }
  deleteBezier(bezierdb: BezierAPI | number, GONG__StackPath: string): Observable<BezierAPI> {
    const id = typeof bezierdb === 'number' ? bezierdb : bezierdb.ID;
    const url = `${this.beziersUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<BezierAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted bezierdb id=${id}`)),
      catchError(this.handleError<BezierAPI>('deleteBezier'))
    );
  }

  // updateFront copy bezier to a version with encoded pointers and update to the back
  updateFront(bezier: Bezier, GONG__StackPath: string): Observable<BezierAPI> {
    let bezierAPI = new BezierAPI
    CopyBezierToBezierAPI(bezier, bezierAPI)
    const id = typeof bezierAPI === 'number' ? bezierAPI : bezierAPI.ID
    const url = `${this.beziersUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<BezierAPI>(url, bezierAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<BezierAPI>('updateBezier'))
    );
  }

  /** PUT: update the bezierdb on the server */
  update(bezierdb: BezierAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<BezierAPI> {
    return this.updateBezier(bezierdb, GONG__StackPath, frontRepo)
  }
  updateBezier(bezierdb: BezierAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<BezierAPI> {
    const id = typeof bezierdb === 'number' ? bezierdb : bezierdb.ID;
    const url = `${this.beziersUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<BezierAPI>(url, bezierdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated bezierdb id=${bezierdb.ID}`)
      }),
      catchError(this.handleError<BezierAPI>('updateBezier'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in BezierService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("BezierService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}
