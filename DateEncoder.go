// DateEncoder.go
//  ----------------------------------------------------------------------
//  wut izit implementation and framework for Hierarchical Temporal Memory
// 
//  Copyright (C) 2018, Charles Perkins.  Unless you have an agreement
//  with Charles Perkins for a separate license for this software code, the
//  following terms and conditions apply:
// 
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU Affero Public License version 3 as
//  published by the Free Software Foundation.
// 
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
//  See the GNU Affero Public License for more details.
// 
//  You should have received a copy of the GNU Affero Public License
//  along with this program.  If not, see http://www.gnu.org/licenses.
// 
//  ----------------------------------------------------------------------
// 


package wut

type DateEncoder struct {
    w int64
    n int64
    season int64
    dayOfWeek int64
    weekend int64
    holiday int64
    timeOfDay int64
    customDays int64
    name string
    forced bool
    sE ScalarEncoder
    sEo int64
    dE ScalarEncoder
    dEo int64
    wE ScalarEncoder
    wEo int64
    hE ScalarEncoder
    hEo int64
    tE ScalarEncoder
    tEo int64
    cE ScalarEncoder
    cEo int64
}


