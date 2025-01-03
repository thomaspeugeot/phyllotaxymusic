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

import { ParameterAPI } from './parameter-api'
import { Parameter, CopyParameterToParameterAPI } from './parameter'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { RhombusAPI } from './rhombus-api'
import { CircleAPI } from './circle-api'
import { RhombusGridAPI } from './rhombusgrid-api'
import { CircleGridAPI } from './circlegrid-api'
import { AxisAPI } from './axis-api'
import { AxisGridAPI } from './axisgrid-api'
import { BezierAPI } from './bezier-api'
import { BezierGridAPI } from './beziergrid-api'
import { BezierGridStackAPI } from './beziergridstack-api'
import { SpiralRhombusAPI } from './spiralrhombus-api'
import { SpiralRhombusGridAPI } from './spiralrhombusgrid-api'
import { SpiralCircleAPI } from './spiralcircle-api'
import { SpiralCircleGridAPI } from './spiralcirclegrid-api'
import { SpiralLineAPI } from './spiralline-api'
import { SpiralLineGridAPI } from './spirallinegrid-api'
import { SpiralBezierAPI } from './spiralbezier-api'
import { SpiralBezierGridAPI } from './spiralbeziergrid-api'
import { FrontCurveStackAPI } from './frontcurvestack-api'
import { KeyAPI } from './key-api'
import { HorizontalAxisAPI } from './horizontalaxis-api'
import { VerticalAxisAPI } from './verticalaxis-api'
import { SpiralOriginAPI } from './spiralorigin-api'

@Injectable({
  providedIn: 'root'
})
export class ParameterService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  ParameterServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private parametersUrl: string

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
    this.parametersUrl = origin + '/api/github.com/thomaspeugeot/phyllotaxymusic/go/v1/parameters';
  }

  /** GET parameters from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<ParameterAPI[]> {
    return this.getParameters(GONG__StackPath, frontRepo)
  }
  getParameters(GONG__StackPath: string, frontRepo: FrontRepo): Observable<ParameterAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<ParameterAPI[]>(this.parametersUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<ParameterAPI[]>('getParameters', []))
      );
  }

  /** GET parameter by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ParameterAPI> {
    return this.getParameter(id, GONG__StackPath, frontRepo)
  }
  getParameter(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ParameterAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.parametersUrl}/${id}`;
    return this.http.get<ParameterAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched parameter id=${id}`)),
      catchError(this.handleError<ParameterAPI>(`getParameter id=${id}`))
    );
  }

  // postFront copy parameter to a version with encoded pointers and post to the back
  postFront(parameter: Parameter, GONG__StackPath: string): Observable<ParameterAPI> {
    let parameterAPI = new ParameterAPI
    CopyParameterToParameterAPI(parameter, parameterAPI)
    const id = typeof parameterAPI === 'number' ? parameterAPI : parameterAPI.ID
    const url = `${this.parametersUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<ParameterAPI>(url, parameterAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<ParameterAPI>('postParameter'))
    );
  }
  
  /** POST: add a new parameter to the server */
  post(parameterdb: ParameterAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ParameterAPI> {
    return this.postParameter(parameterdb, GONG__StackPath, frontRepo)
  }
  postParameter(parameterdb: ParameterAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ParameterAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<ParameterAPI>(this.parametersUrl, parameterdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted parameterdb id=${parameterdb.ID}`)
      }),
      catchError(this.handleError<ParameterAPI>('postParameter'))
    );
  }

  /** DELETE: delete the parameterdb from the server */
  delete(parameterdb: ParameterAPI | number, GONG__StackPath: string): Observable<ParameterAPI> {
    return this.deleteParameter(parameterdb, GONG__StackPath)
  }
  deleteParameter(parameterdb: ParameterAPI | number, GONG__StackPath: string): Observable<ParameterAPI> {
    const id = typeof parameterdb === 'number' ? parameterdb : parameterdb.ID;
    const url = `${this.parametersUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<ParameterAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted parameterdb id=${id}`)),
      catchError(this.handleError<ParameterAPI>('deleteParameter'))
    );
  }

  // updateFront copy parameter to a version with encoded pointers and update to the back
  updateFront(parameter: Parameter, GONG__StackPath: string): Observable<ParameterAPI> {
    let parameterAPI = new ParameterAPI
    CopyParameterToParameterAPI(parameter, parameterAPI)
    const id = typeof parameterAPI === 'number' ? parameterAPI : parameterAPI.ID
    const url = `${this.parametersUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<ParameterAPI>(url, parameterAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<ParameterAPI>('updateParameter'))
    );
  }

  /** PUT: update the parameterdb on the server */
  update(parameterdb: ParameterAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ParameterAPI> {
    return this.updateParameter(parameterdb, GONG__StackPath, frontRepo)
  }
  updateParameter(parameterdb: ParameterAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<ParameterAPI> {
    const id = typeof parameterdb === 'number' ? parameterdb : parameterdb.ID;
    const url = `${this.parametersUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<ParameterAPI>(url, parameterdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated parameterdb id=${parameterdb.ID}`)
      }),
      catchError(this.handleError<ParameterAPI>('updateParameter'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in ParameterService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("ParameterService" + error); // log to console instead

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
